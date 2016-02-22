# 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

# What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

function divisible_upto(num, n)
  counter = 1
  while counter <= n
    if mod(num, counter) != 0
      return false
    end
    counter += 1
  end
  true
end

function divisible_by_first_20()
  result = 1000
  while true
    if divisible_upto(result, 20)
      return result
    end
    result += 1
  end
end

println(divisible_by_first_20())
