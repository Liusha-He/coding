import torch


if __name__ == '__main__':
    x = torch.randn(3, requires_grad=True)
    print(f"{x=}")

    y = x + 2
    print(f"{y=}")

    z = y * y * 2
    z = z.mean()
    print(f"{z=}")

    z.backward()
    print(x.grad)
