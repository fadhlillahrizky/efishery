const { Sequelize } = require('sequelize');
const db = require('../config/db');
const { DataTypes } = Sequelize;


const attributes = {
    id: {
        type: DataTypes.BIGINT,
        primaryKey: true,
        autoIncrement: true,
    },
    name: {
        type: DataTypes.STRING,
    },
    phone: {
        type: DataTypes.STRING,
    },
    role: {
        type: DataTypes.STRING,
    },
    password: {
        type: DataTypes.STRING,
    },
    createdAt: {
        type: DataTypes.DATE,
    },
    updatedAt: {
        type: DataTypes.DATE,
    },
    deletedAt: {
        type: DataTypes.DATE,
    },
};

const Users = db.define('users', attributes, {
    freezeTableName: true,
    underscored: true,
});


module.exports = Users;
