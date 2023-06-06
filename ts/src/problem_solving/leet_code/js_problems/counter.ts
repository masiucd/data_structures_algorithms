export const createCounter = function (n: number) {
  return function () {
    return n++;
  };
};
