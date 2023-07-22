type Role = 'ADMIN' | 'USER';

type User = {
  userId: string;
  username: string;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
  role: Role;
};

export { User };
