using Plots

function logistic_bifurcation()
  rs = range(0.0, 4.0, length=2000)
  max_step = 1000
  keep_step = 200
  x0 = 0.5

  x_coords = Float64[]
  y_coords = Float64[]

  for r in rs
    x = x0
    for _ in 1:(max_step-keep_step)
      x = r * x * (1 - x)
    end
    for _ in 1:keep_step
      x = r * x * (1 - x)
      push!(x_coords, r)
      push!(y_coords, x)
    end
  end

  scatter(x_coords, y_coords,
    markersize=0.5,
    color=:black,
    xlabel="r",
    ylabel="x",
    title="Logistic Map Bifurcation Diagram",
    xlim=(0, 4),
    ylim=(0, 1),
    legend=false,
    size=(2000, 1000)
  )

  savefig("bifurcation_diagram.png")
end

logistic_bifurcation()
