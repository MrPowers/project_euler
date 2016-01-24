def palindrome?(n)
  str = n.to_s
  return false if str[0] == "0"
  str == str.reverse
end

result = (1..999_999).to_a.inject(0) do |sum, n|
  if palindrome?(n) && palindrome?(n.to_s(2))
    sum += n
  end
  sum
end

p result
