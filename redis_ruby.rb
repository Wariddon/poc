require 'redis'
require 'benchmark'

redis = Redis.new(
  :host => 'localhost',
  :port => 6379,
  :db => 7
)

N = 1000
time = Benchmark.realtime do
  N.times do
    value = redis.get("test_key")
  end
end

puts "Ruby Redis Total Time for #{N} fetches from DB 7: #{time} seconds"