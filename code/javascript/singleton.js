import exampleDB from "exampledb";

export class DataBase {
  constructor() {
    this.connection = exampleDB();
  }
  getLink(id) {
    return this.connection.get(id);
  }

  deleteByID(id) {
    return connection.delete(id);
  }
}

export function generateDB() {
  if (!generateDB.instance) {
    generateDB.instance = {
      connection: exampleDB(),

      getLink(id) {
        return this.connection.get(id);
      },

      deleteByID(id) {
        this.connection.delete(id);
      },
    };
  }

  return generateDB.instance;
}

export default new DataBase();
