using estuary_backend.Models;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;

namespace estuary_backend
{
    public class EstuaryDbContext : DbContext
    {
        protected override void OnConfiguring(DbContextOptionsBuilder options) => options.UseSqlite("Data Source=estuary_backend.db");

        public DbSet<Workshop> Workshops { get; set; }

        public DbSet<PropblemStatement> PropblemStatements { get; set; }

        public DbSet<Tag> Tags { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            #region seed test data

            var firstTag = new Tag { Name = "Improvements", Id = Guid.NewGuid() };

            modelBuilder.Entity<Tag>().HasData(firstTag);


            var firstWorkshop = new Workshop
            {
                Id = Guid.NewGuid(),
                Date = DateTime.Now,
                Authors = new List<Author> { },
                Teaser = "Lorem ipsum dolor sit amet.",
                LocationName = "Hamburg", 
                LocationMap = "https://goo.gl/maps/mbnen1jr8C81J6vU9"
            };
            modelBuilder.Entity<Workshop>().HasData(firstWorkshop);


            var firstProblemStatement = new PropblemStatement
            {
                Id = Guid.NewGuid(),
                WorkshopId = firstWorkshop.Id,
                Title = "ProblemStatement 1",
                Iam = "Lorem ipsum",
                Iwant = "dolor sit amet, consetetur sadipscing elitr",
                But = "sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat",
                Because = "sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.",
                Feel = "Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.Lorem ipsum dolor sit amet, consetetur sadipscing."
            };

            modelBuilder.Entity<PropblemStatement>().HasData(firstProblemStatement);


            var workshopTag = new WorkshopTag { Id = Guid.NewGuid(), TagId = firstTag.Id, WorkshopId = firstWorkshop.Id };

            modelBuilder.Entity<WorkshopTag>().HasData(workshopTag);


            #endregion

            base.OnModelCreating(modelBuilder);

        }
    }
}
