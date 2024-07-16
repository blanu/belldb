CREATE MIGRATION m17x4oljceelwnrszaxcy7cs56w3yohkbkfzfp2ykswl7qzpjew62q
    ONTO m1guehzs6jlcnayqyzlz3lv6hw753v4ov4kqlxmpmvz2iq7aqew7oa
{
  CREATE TYPE default::Definition {
      CREATE LINK `module`: default::`Module`;
      CREATE PROPERTY label: std::str;
      CREATE CONSTRAINT std::exclusive ON ((.`module`, .label));
      CREATE PROPERTY expression: std::str;
  };
  ALTER TYPE default::`Module` {
      ALTER LINK dependencies {
          USING (.<`module`);
          RESET CARDINALITY;
      };
  };
};
