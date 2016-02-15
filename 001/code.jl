#If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

#Find the sum of all the multiples of 3 or 5 below 1000.

function getnumbers()
  result = []
  for num in collect(1:999)
    if mod(num, 3) == 0 || mod(num, 5) == 0
      push!(result, num)
    end
  end
  result
end

function sumarray(arr)
  result = 0
  for num in arr
    result += num
  end
  result
end

println(sumarray(getnumbers()))
