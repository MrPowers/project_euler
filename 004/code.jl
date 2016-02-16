# A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

# Find the largest palindrome made from the product of two 3-digit numbers.

function palindromeNumbers(a, b)
  result = []
  for num in collect(a:b)
    s = string(num)
    if s == reverse(s)
      push!(result, s)
    end
  end
  result
end

function isPalindrome(num)
  s = string(num)
  s == reverse(s)
end

function threeDigitProducts()
  result = []
  for i in collect(100:999)
    for j in collect(i:999)
      push!(result, i*j)
    end
  end
  result
end

nums = sort(threeDigitProducts())
index = findlast(x->isPalindrome(x), nums)
println(nums[index])
