import torch
from torch.utils.data import DataLoader
import torch.nn as nn
import torchvision
import torchvision.transforms as transforms
import matplotlib.pyplot as plt


class NeuralNet(nn.Module):
    def __init__(self, input_size: int, hidden_size: int, num_classes: int):
        super(NeuralNet, self).__init__()
        self.l1 = nn.Linear(input_size, hidden_size)
        self.relu = nn.ReLU()
        self.l2 = nn.Linear(hidden_size, num_classes)

    def forward(self, x):
        out = self.l1(x)
        out = self.relu(out)
        out = self.l2(out)
        return out


if __name__ == '__main__':
    model_path = "../models/digit_recognition/model.pkl"
    device = torch.device("cuda" if torch.cuda.is_available() else "cpe")

    input_size = 784  # 28 x 28
    hidden_size = 100
    num_classes = 10
    num_epochs = 5
    batch_size = 100
    learning_rate = .001

    train_dataset = torchvision.datasets.MNIST(root='../data/mnist',
                                               train=True,
                                               transform=transforms.ToTensor(),
                                               download=True)
    test_dataset = torchvision.datasets.MNIST(root='../data/mnist',
                                              train=False,
                                              transform=transforms.ToTensor(),
                                              download=False)

    train_loader = DataLoader(dataset=train_dataset,
                              batch_size=batch_size,
                              shuffle=True)
    test_loader = DataLoader(dataset=test_dataset,
                             batch_size=batch_size,
                             shuffle=False)

    # model = NeuralNet(input_size, hidden_size, num_classes)
    # model.to(device)
    #
    # loss_fn = nn.CrossEntropyLoss()
    # optimizer = torch.optim.Adam(model.parameters(), lr=learning_rate)
    #
    # n_total_steps = len(train_loader)
    #
    # for epoch in range(num_epochs):
    #     for i, (images, labels) in enumerate(train_loader):
    #         images = images.reshape(-1, 784).to(device)
    #         labels = labels.to(device)
    #
    #         outputs = model(images)
    #         loss = loss_fn(outputs, labels)
    #
    #         optimizer.zero_grad()
    #         loss.backward()
    #         optimizer.step()
    #
    #         if (i + 1) % 100 == 0:
    #             print(f"epoch {epoch + 1} / {num_epochs}: step - {i + 1} / {n_total_steps} : loss = {loss.item():.4f}")

    # # save model
    # torch.save(model.state_dict(), model_path)

    # load model
    model = NeuralNet(input_size, hidden_size, num_classes)
    model.load_state_dict(torch.load(model_path))
    model.eval()
    model.to(device)

    # test
    with torch.no_grad():
        n_correct = 0
        n_samples = 0

        for images, labels in test_loader:
            images = images.reshape(-1, 784).to(device)
            labels = labels.to(device)

            outputs = model(images)

            _, predictions = torch.max(outputs, 1)
            n_samples += labels.shape[0]
            n_correct += (predictions == labels).sum().item()

        acc = 100.0 * n_correct / n_samples
        print(f"Accuracy = {acc}%")

        examples = iter(test_loader)
        samples, labels = next(examples)

        for i in range(12):
            plt.subplot(3, 4, i + 1)

            idx = i + 20

            image = samples[idx].reshape(-1, 784).to(device)
            l = labels[idx].item()

            out = model(image)
            _, pred = torch.max(out, 1)
            plt.imshow(samples[idx][0], cmap="gray")
            plt.title(f"{pred.item()} (actual {l})")
            plt.axis("off")
        plt.show()
