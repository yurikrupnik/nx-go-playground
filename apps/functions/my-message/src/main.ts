import type { HttpFunction } from '@google-cloud/functions-framework/build/src/functions';

const { foo = 'testBar' } = process.env;

export const MyMessage: HttpFunction = async (req, res) => {
  res.status(200).send(foo);
};
