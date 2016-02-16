# The prime factors of 13195 are 5, 7, 13 and 29.

# What is the largest prime factor of the number 600851475143 ?

function firstPrimeFactor(n)
  num = 2
  while num <= n
    if mod(n, num) == 0
      return num
    end
    num += 1
  end
end

function primeFactors(n)
  result = []
  while true
    p = firstPrimeFactor(n)
    push!(result, p)
    if p == n
      return result
    end
    n = div(n, p)
  end
end

function maxValue(arr)
  result = arr[1]
  for num in arr
    if num > result
      result = num
    end
  end
  result
end

println(maxValue(primeFactors(600851475143)))
