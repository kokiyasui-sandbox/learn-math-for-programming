using Plots

logistic_map(r, x) = r * x * (1 - x)

function trajectory(r, x0, step)
  trajectories = zeros(step)
  trajectories[1] = x0

  for i in 1:(step-1)
    trajectories[i+1] = logistic_map(r, trajectories[i])
  end

  return trajectories
end

function main()
  rs = [0.5, 2.4, 3.3, 3.5, 3.5644072661, 3.9]
  x0 = 0.5
  max_step = 1000
  plot_start=900

  plots = Plots.Plot[]

  for r in rs
    t = trajectory(r, x0, max_step)
    p = plot(
      plot_start:max_step,
      t[plot_start:max_step],
      title="r = $r",
      xlabel="n",
      ylabel="x_n",
      ylim=(0, 1),
      legend=false
    )
    push!(plots, p)
  end

  final_plot = plot(plots..., layout=(3, 2), size=(2000, 1000))

  savefig(final_plot, "logistic_trajectories.png")
end

# 修正3: End -> end
if abspath(PROGRAM_FILE) == @__FILE__
  main()
end
