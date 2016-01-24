require_relative "names"

def mapping
  @mapping ||= ("A".."Z").to_a.each_with_index.inject({}) do |memo, (letter, i)|
    memo[letter] = i + 1
    memo
  end
end

def to_number(name)
  name.split("").inject(0) do |sum, letter|
    sum += mapping[letter]
    sum
  end
end

result = NAMES.sort.each_with_index.inject(0) do |sum, (name, index)|
  sum += to_number(name) * (index + 1)
  sum
end

p result
