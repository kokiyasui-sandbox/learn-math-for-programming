fibonacci(n) = round(Int, 1 / sqrt(5) * ((1 + sqrt(5))/2) ^ n - 1 / sqrt(5) * ((1 - sqrt(5))/2) ^ n)

function main()
  for i in 1:20
    println(fibonacci(i))
  end
end

main()
