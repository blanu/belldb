module default {
  type `Module` {
    required name: str {
      constraint exclusive;
    }

    multi dependencies: `Module`;

    definitions := .<`module`[is Definition];
  }

  type Definition {
    required `module`: `Module`;
    required label: str;
    expression: str;

    constraint exclusive on ((.`module`, .label));
  }
}
