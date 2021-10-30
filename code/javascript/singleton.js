class DataBase {
  constructor() {
    DBDriver.connect();
  }

  async getLink(id) {
    // Do stuff
  }

  async deleteByID(id) {
    // Do other stuff
  }
}

const DataBaseTwo = {
  getLink: (id) => {
    // Do stuff
  },

  deleteByID: (id) => {
    // Do other stuff
  },
};

export default new DataBase();
