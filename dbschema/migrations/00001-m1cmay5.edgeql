CREATE MIGRATION m1cmay5gpfn63jbe52oq6dofs6whhuzjrq4kvgbvnbhg7q6x7kff3q
    ONTO initial
{
  CREATE TYPE default::Definition {
      CREATE PROPERTY expression: std::str;
      CREATE PROPERTY label: std::str;
  };
  CREATE TYPE default::`Module` {
      CREATE MULTI LINK definitions: default::Definition;
      CREATE MULTI LINK dependencies: default::`Module`;
      CREATE PROPERTY name: std::str;
  };
};
