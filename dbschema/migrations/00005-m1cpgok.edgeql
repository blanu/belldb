CREATE MIGRATION m1cpgoksnguxl5e5xqyahczl3x3nmfkbsvxcvf32rucc42kugrevhq
    ONTO m1medsknyximk3mirhcfivnlcb6gidw2ipzsjahvycradvdqgjvwxq
{
  ALTER TYPE default::`Module` {
      CREATE MULTI LINK dependencies: default::`Module`;
  };
};
