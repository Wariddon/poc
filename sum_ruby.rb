require 'benchmark'

N = 10_000_000

time = Benchmark.realtime do
  sum = (1..N).reduce(0, :+)
  puts "Sum: #{sum}"
end

puts "Ruby Time: #{time} seconds"
