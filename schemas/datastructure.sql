create table characters (
    id serial,
    name text not null,
    race text not null,
    class text not null,
    base_damage int,
    hp int,
    background text,
    src_image text,
    constraint characters_pk primary key (id)
);

create table powers (
    id serial,
    name text,
    description text not null,
    damage int,
    constraint powers_pk primary key (id)
);

create table rel_char_pow(
    characters_id int,
    powers_id int,
    constraint rel_char_pow_pk primary key (characters_id, powers_id),
    constraint rel_char_pow_characters_fk FOREIGN key (characters_id) references characters(id),
    constraint rel_char_pow_powers_fk FOREIGN key (powers_id) references powers(id)
);