CREATE MIGRATION m1n3x3n3o44u3jfl77x6h7azl4qp4vy2svdd3qk7ddxl44nxaxu6gq
    ONTO m1cpgoksnguxl5e5xqyahczl3x3nmfkbsvxcvf32rucc42kugrevhq
{
  CREATE TYPE default::Builtin {
      CREATE LINK `module`: default::`Module`;
      CREATE PROPERTY label: std::str;
      CREATE CONSTRAINT std::exclusive ON ((.`module`, .label));
  };
};
