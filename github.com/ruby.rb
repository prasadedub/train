module RIO
  module Service
    class DivisionByMac < RIO::Service::RPCServer
      workers 4
      qualifier :market
      uri = "jdbc:mysql://%s:%d/%s?user=%s&password=%s" % [ENV["NYROC_MYSQL_HOST"], ENV["NYROC_MYSQL_PORT"], ENV["NYROC_MYSQL_DB"], ENV["NYROC_MYSQL_USER"], ENV["NYROC_MYSQL_PASS"]]
      DB = Sequel.connect(uri)

      def db
        self.class::DB
      end
      def serve(props, data)
        uptick :requests
        results = Hash.new {|h,k| h[k] = Hash.new }
        Array(data["mac"]).map do |mac|
          sql = "SELECT division FROM mactovendor WHERE macAddress = ?"
          begin
            if row = db.fetch(sql, mac).first
              results[mac][:division] = row[:division]
            else
                results[mac][:error] = 'No Data Present'
            end
          rescue Exception => e
            results[mac] = nil
            puts "Error for ip #{mac}: #{e.message}"
          end
        end
        pp results
        results
      end
    end
  end
end

module RIO
  module Service
    class CmtsByIP < RIO::Service::RPCServer
      workers 4
      qualifier :market
      uri = "jdbc:mysql://%s:%d/%s?user=%s&password=%s" % [ENV["NYROC_MYSQL_HOST"], ENV["NYROC_MYSQL_PORT"], ENV["NYROC_MYSQL_DB"], ENV["NYROC_MYSQL_USER"], ENV["NYROC_MYSQL_PASS"]]
      DB = Sequel.connect(uri)

      def db
        self.class::DB
      end


      def echo_prefix
        "i got"
      end

      def ipv4_range_sql(ip)
        ip = ip.to_i
        sql = <<-THE_END
SELECT device_name, division 
FROM nyroctools.mrtg mrtg
JOIN reporting.`tbDeviceToNetwork` net ON net.deviceName = mrtg.device_name AND mrtg.int_type IN (127,24,142)
WHERE ipv4Network <= #{ip} AND #{ip} <= ipv4Broadcast
LIMIT 1
        THE_END
      end

      def ipv6_range_sql(ip)
        left  = left_side(ip).to_i
        right = right_side(ip).to_i
        sql = <<-THE_END
          SELECT device_name, division
          FROM nyroctools.mrtg mrtg
          JOIN reporting.`tbDeviceToNetwork` net ON net.deviceName = mrtg.device_name AND mrtg.int_type IN (127,24,142)
          WHERE  ipv6RightNetwork <= #{right} AND #{right} <= ipv6RightBroadcast  AND ipv6LeftNetwork = #{left} AND ipv6LeftBroadCast = #{left}
          LIMIT 1
        THE_END
      end

      def left_mask
        @left_mask  ||= IPAddr.new('FFFF:FFFF:FFFF:FFFF::')
      end

      def right_mask
        @right_mask ||= IPAddr.new('::FFFF:FFFF:FFFF:FFFF')
      end

      def right_side(ip)
        right = ip & right_mask
      end

      def left_side(ip)
        left = ip & left_mask
        left >> (64)
      end

      def serve(props, data)
        uptick :requests
        results = Hash.new {|h,k| h[k] = Hash.new }
        Array(data["ip"]).map do |ip|
          begin
            ip   = IPAddr.new(ip)
            if ip.ipv4?
              sql = ipv4_range_sql(ip)
            elsif ip.ipv6?
              sql = ipv6_range_sql(ip)
            end
            if row = db[sql].first
              results[ip][:cmts] = row[:device_name]
              results[ip][:division] = row[:division]
              results[:mac][:error] = nil
            else
                results[ip][:error] = 'No Data Present'
            end
          rescue Exception => e
            results[ip] = {'cmts': nil, 'divisions': nil, 'error': 'No data Present'}
            puts "Error for ip #{ip}: #{e.message}"
          end
        end
        results
      end

    end
  end
end