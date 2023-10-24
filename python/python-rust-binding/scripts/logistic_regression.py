from datetime import datetime

import torch
import torch.nn as nn
import numpy as np
from sklearn import datasets
from sklearn.preprocessing import StandardScaler
from sklearn.model_selection import train_test_split


class LogisticRegression(nn.Module):
    def __init__(self, n_input_features):
        super(LogisticRegression, self).__init__()
        self.linear = nn.Linear(n_input_features, 1)

    def forward(self, x):
        y_pred = torch.sigmoid(self.linear(x))
        return y_pred


if __name__ == '__main__':
    bc = datasets.load_breast_cancer()
    x, y = bc.data, bc.target

    n_sample, n_features = x.shape

    x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=.2, random_state=42)

    sc = StandardScaler()
    x_train = sc.fit_transform(x_train)
    x_test = sc.transform(x_test)

    # convert data to tensor and move to gpu
    x_train = torch.from_numpy(x_train.astype(np.float32))
    x_test = torch.from_numpy(x_test.astype(np.float32))
    y_train = torch.from_numpy(y_train.astype(np.float32)).reshape(y_train.shape[0], 1)
    y_test = torch.from_numpy(y_test.astype(np.float32)).reshape(y_test.shape[0], 1)

    # x_train = x_train.to("cuda")
    # x_test = x_test.to("cuda")
    # y_train = y_train.to("cuda")
    # y_test = y_test.to("cuda")

    lr = .01

    # create model and move to gpu
    model = LogisticRegression(n_features)
    # model.to("cuda")

    loss_fn = nn.BCELoss()

    optimizer = torch.optim.Adam(model.parameters(), lr=lr)
    n_epochs = 100

    for epoch in range(n_epochs):
        y_pred = model(x_train)

        l = loss_fn(y_pred, y_train)

        optimizer.zero_grad()
        l.backward()

        optimizer.step()

        if epoch % 5 == 0:
            print(f"epoch {epoch+1} => loss = {l.item():.4f}.")

    start = datetime.now()
    with torch.no_grad():
        preds = model(x_test)
        print(f"The prediction took {(datetime.now() - start).microseconds}")

        preds_cls = preds.round()

        acc = preds_cls.eq(y_test).sum() / float(y_test.shape[0])
        print(f"accuracy = {acc:.4f}")


