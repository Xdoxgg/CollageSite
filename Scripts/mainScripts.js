function getItemFromLocalStorage(key) {
    return new Promise((resolve, reject) => {
      try {
        const item = localStorage.getItem(key);
        resolve(item);
      } catch (error) {
        reject(error);
      }
    });
  }