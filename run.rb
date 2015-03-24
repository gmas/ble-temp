#!/bin/env ruby
require 'date'
require 'json'

IO.popen('./test.js').each_line.grep(/received:/) do |line|
  #puts "MATCH: #{line}"
  json_data = line.split('received:')[1].strip
  data = JSON.parse json_data
  p data
  #rec_date = DateTime.strptime(data['date'], '%Y-%m-%dT%H:%M:%S')
  #p rec_date
end
