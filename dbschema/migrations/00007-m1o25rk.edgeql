CREATE MIGRATION m1o25rkpgghbcbg6ncfd36wbz4ud7j32gx427qidtdn2f4365lfdpq
    ONTO m1n3x3n3o44u3jfl77x6h7azl4qp4vy2svdd3qk7ddxl44nxaxu6gq
{
  ALTER TYPE default::`Module` {
      CREATE LINK builtins := (.<`module`[IS default::Builtin]);
      ALTER LINK definitions {
          USING (.<`module`[IS default::Definition]);
      };
  };
};
