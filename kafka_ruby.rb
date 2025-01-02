require 'kafka'
require 'benchmark'

kafka = Kafka.new(["127.0.0.1:9092"])
N = 1000 # Number of messages to produce and consume

# Measure time for producing messages
produce_time = Benchmark.realtime do
  producer = kafka.producer
  N.times do |i|
    producer.produce("Message #{i}", topic: "test_topic")
  end
  producer.deliver_messages
end
puts "Ruby - Produced #{N} messages in #{produce_time.round(4)} seconds."

# Measure time for consuming messages
consume_time = Benchmark.realtime do
  consumer = kafka.consumer(group_id: "test_group")
  consumer.subscribe("test_topic")

  count = 0
  consumer.each_message do |message|
    count += 1
    break if count >= N
  end
end
puts "Ruby - Consumed #{N} messages in #{consume_time.round(4)} seconds."