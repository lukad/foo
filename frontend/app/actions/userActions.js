export function setAge(age) {
  return {
    type: "SET_AGE",
    payload: age
  };
};

export function setName(name) {
  return {
    type: "SET_NAME",
    payload: new Promise((resolve, _reject) => {
      setTimeout(() => {
        resolve(name);
      }, 2000);
    })
  };
};
