from os import path

import matplotlib.pyplot as plt
import pandas as pd


def plot_quantiles(netlogo_file, beecs_file, out_dir, format):
    data_beehave = pd.read_csv(netlogo_file, delimiter=";")
    data_beecs = pd.read_csv(beecs_file, delimiter=";")

    columns = list(data_beehave.columns)[1:]
    columns = pd.unique([c[:-4] for c in columns])
    quantiles = [
        ("Q05", 5),
        ("Q10", 10),
        ("Q25", 25),
        ("Q50", 50),
        ("Q75", 75),
        ("Q90", 90),
        ("Q95", 95),
    ]

    for col in columns:
        plot_column(
            data_beehave,
            data_beecs,
            col,
            quantiles,
            path.join(out_dir, col + "." + format),
        )


def plot_column(data_beehave, data_beecs, column, quantiles, image_file):
    median_col = quantiles[len(quantiles) // 2][0]

    fig, ax = plt.subplots(figsize=(10, 4))
    for data, col, model in [
        (data_beehave, "red", "BEEHAVE"),
        (data_beecs, "blue", "beecs"),
    ]:
        q50 = data[column + "_Q50"]
        q10 = data[column + "_Q05"]
        q90 = data[column + "_Q95"]

        ax.plot(data.ticks, q50, c=col, label=model)
        ax.fill_between(data.ticks, q10, q90, color=col, alpha=0.1)

    ax.set_title(column)
    ax.set_xlabel("time [d]", fontsize="12")
    ax.legend()
    fig.tight_layout()

    plt.savefig(image_file)
    plt.close()


if __name__ == "__main__":
    plot_quantiles(
        "_tests/default/beehave.csv",
        "_tests/default/beecs.csv",
        "_tests/default",
        #"png",
        "svg",
    )
