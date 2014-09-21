file = File.expand_path("./numbers", File.dirname(__FILE__))
numbers = File.readlines(file).map(&:chomp)
sum = numbers.inject(0) do |memo, number|
  memo += number.to_i
  memo
end
p sum.to_s[0, 10]
