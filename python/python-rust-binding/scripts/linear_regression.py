import torch
import torch.nn as nn
import numpy as np
from sklearn import datasets
import matplotlib.pyplot as plt


if __name__ == '__main__':
    x_numpy, y_numpy = datasets.make_regression(
        n_samples=100,
        n_features=1,
        noise=20,
        random_state=42
    )

    x = torch.from_numpy(x_numpy.astype(np.float32))
    y = torch.from_numpy(y_numpy.astype(np.float32))
    y = y.view(y.shape[0], 1)

    x_test = torch.tensor([5], dtype=torch.float32)
    n_samples, n_features = x.shape
    input_size = n_features
    output_size = n_features

    model = nn.Linear(input_size, output_size)

    lr = 0.01
    loss = nn.MSELoss()
    optimizer = torch.optim.SGD(model.parameters(), lr=lr)
    n_iters = 100

    for epoch in range(n_iters):
        y_pred = model(x)
        l = loss(y, y_pred)
        l.backward()

        optimizer.step()
        optimizer.zero_grad()

        if epoch % 2 == 0:
            (w, b) = model.parameters()
            print(f"epoch {epoch - 1} => loss = {l:.3f}.")

    predictions = model(x).detach()

    plt.plot(x_numpy, y_numpy, "ro")
    plt.plot(x_numpy, predictions, "b")
    plt.show()

    torch.save(model.state_dict(), "linear.pkl")
