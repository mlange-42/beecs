from os import path

import numpy as np
import pandas as pd


def agg_beecs(file_pattern, out_file):
    data = None

    idx = 0
    while True:
        file = file_pattern % (idx,)
        if not path.exists(file):
            break

        run = pd.read_csv(file, delimiter=";")
        run = run.rename(columns={"t": "ticks"})
        run.insert(1, "run", idx)
        if data is None:
            data = run
        else:
            data = pd.concat([data, run])

        idx += 1

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
            out[col] = 0.
        for tick in ticks:
            values = data[column][data.ticks == tick]
            q = np.quantile(values, [0.05, 0.1, 0.25, 0.5, 0.75, 0.9, 0.95])
            out.loc[tick, cols] = q

    out.to_csv(out_file, sep=";", index=False)


if __name__ == "__main__":
    agg_beecs("_tests/default/out/beecs-%04d.csv", "_tests/default/beecs.csv")
