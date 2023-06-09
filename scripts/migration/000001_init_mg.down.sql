alter table classes drop foreign key classes_ibfk_3;

drop table if exists enrolls cascade;

drop table if exists rooms cascade;

drop table if exists classes cascade;

drop table if exists courses cascade;

drop table if exists professors cascade;

drop table if exists students cascade;