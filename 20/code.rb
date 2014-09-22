def factorial(n)
  (1..n).to_a.inject(&:*)
end

def sum_digits(n)
  n.to_s.split("").inject(0) do |memo, n|
    memo += n.to_i
    memo
  end
end

p sum_digits(factorial(100))
