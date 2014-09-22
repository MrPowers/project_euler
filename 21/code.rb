#Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
#If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

#For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

#Evaluate the sum of all the amicable numbers under 10000.
#
#

def proper_divisors(n)
  (1..n-1).to_a.inject([]) do |memo, num|
    memo << num if n % num == 0
    memo
  end
end

def d(n)
  proper_divisors(n).inject(&:+)
end

def mapping(first, last)
  (first..last).to_a.inject({}) do |memo, n|
    memo[n] = d(n)
    memo
  end
end

@mapping = mapping(1, 9999)

def amicable_pairs
  result = []
  @mapping.each do |k, d|
    if k == @mapping[d] && k != d && !result.include?([k, d].sort)
      result << [k, d]
    end
  end
  result
end

p amicable_pairs.flatten.inject(&:+)
