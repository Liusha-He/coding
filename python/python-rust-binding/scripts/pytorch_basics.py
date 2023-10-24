import torch
import numpy as np


if __name__ == '__main__':
    array = np.array([1,2,3,4,5,6], dtype=np.float64)
    tensor = torch.from_numpy(array).reshape([2,3])
    print(tensor)

    if torch.cuda.is_available():
        device = torch.device("cuda")
        x = torch.tensor([[1,2,3],[4,5,6]],
                         device=device)
        y = torch.tensor([[1,1,1],[2,2,2]])
        y = y.to(device)

        z = x + y
        print(z)

        # if comment out below, it will throw an error, coz numpy work on cpu.
        # z.numpy()

        z = z.to("cpu")
        print(z.numpy())
