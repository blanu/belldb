CREATE MIGRATION m1guehzs6jlcnayqyzlz3lv6hw753v4ov4kqlxmpmvz2iq7aqew7oa
    ONTO m1cmay5gpfn63jbe52oq6dofs6whhuzjrq4kvgbvnbhg7q6x7kff3q
{
  ALTER TYPE default::`Module` {
      DROP LINK definitions;
  };
  ALTER TYPE default::`Module` {
      CREATE PROPERTY definitions: array<tuple<std::str, std::str>>;
      ALTER PROPERTY name {
          CREATE CONSTRAINT std::exclusive;
          SET REQUIRED USING ('default');
      };
  };
  DROP TYPE default::Definition;
};
