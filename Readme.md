golang movies project

1)Implemented apis for getting data related to movies

2)Used local mysql database and created the database movies and table as movies

create table `movies` (
	`title` varchar (300),
	`released_year` varchar (300),
	`rating` varchar (300),
	`id` varchar (300),
	`generes` varchar (300)
);
insert into `movies` (`title`, `released_year`, `rating`, `id`, `generes`) values('Krish','2010','3','1','Crime');
insert into `movies` (`title`, `released_year`, `rating`, `id`, `generes`) values('Sultan ','2010','2','2','Sport');

