CREATE MIGRATION m1medsknyximk3mirhcfivnlcb6gidw2ipzsjahvycradvdqgjvwxq
    ONTO m17x4oljceelwnrszaxcy7cs56w3yohkbkfzfp2ykswl7qzpjew62q
{
  ALTER TYPE default::`Module` {
      DROP PROPERTY definitions;
  };
  ALTER TYPE default::`Module` {
      CREATE LINK definitions := (.<`module`);
      DROP LINK dependencies;
  };
};
