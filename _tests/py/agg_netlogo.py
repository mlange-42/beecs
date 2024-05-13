import numpy as np
import pandas as pd


def agg_netlogo(file, out_file):
    data = pd.read_csv(file, delimiter=";")

    runs = pd.unique(data.run)
    runs.sort()
    ticks = pd.unique(data.ticks)
    ticks.sort()

    columns = list(data.columns)[2:]

    out = pd.DataFrame(data={"ticks": ticks}, index=ticks)

    for column in columns:
        cols = [
            column + "_Q05",
            column + "_Q10",
            column + "_Q25",
            column + "_Q50",
            column + "_Q75",
            column + "_Q90",
            column + "_Q95",
        ]
        for col in cols:
            out[col] = 0
        for tick in ticks:
            values = data[column][data.ticks == tick]
            q = np.quantile(values, [0.05, 0.1, 0.25, 0.5, 0.75, 0.9, 0.95])
            out.loc[tick, cols] = q

    out.to_csv(out_file, sep=";", index=False)


if __name__ == "__main__":
    agg_netlogo("out/netlogo.csv", "tests/default/netlogo.csv")
