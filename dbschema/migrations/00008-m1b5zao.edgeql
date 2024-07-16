CREATE MIGRATION m1b5zaoa5lblx7lo2x6p7cyb5buakclxlajbl26cqo5w5ndwumyipa
    ONTO m1o25rkpgghbcbg6ncfd36wbz4ud7j32gx427qidtdn2f4365lfdpq
{
  ALTER TYPE default::Builtin {
      DROP CONSTRAINT std::exclusive ON ((.`module`, .label));
  };
  ALTER TYPE default::`Module` {
      DROP LINK builtins;
  };
  DROP TYPE default::Builtin;
  ALTER TYPE default::Definition {
      ALTER LINK `module` {
          SET REQUIRED USING (<default::`Module`>{});
      };
      ALTER PROPERTY label {
          SET REQUIRED USING (<std::str>{});
      };
  };
};
