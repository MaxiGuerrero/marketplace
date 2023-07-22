type AddProductRequest = {
  productId: string;
  amount: number;
};

type RemoveProductRequest = {
  productId: string;
};

export { AddProductRequest, RemoveProductRequest };
