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

            var firstTag = new Tag { Name = "Improvements" };

            var firstProblemStatement = new PropblemStatement
            {
                Id = Guid.NewGuid(),
                Title = "ProblemStatement 1",
                Iam = "Lorem ipsum",
                Iwant = "dolor sit amet, consetetur sadipscing elitr",
                But = "sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat",
                Because = "sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.",
                Feel = "Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.Lorem ipsum dolor sit amet, consetetur sadipscing."
            };

            var firstWorkshop = new Workshop
            {
                Id = Guid.NewGuid(),
                Date = DateTime.Now,
                Authors = new List<Author> { },
                Teaser = "Lorem ipsum dolor sit amet.",
                Location = new Location { Name = "Hamburg", Map = "https://goo.gl/maps/mbnen1jr8C81J6vU9" },
            };
            firstWorkshop.Tags.Add(firstTag);
            firstTag.Workshops.Add(firstWorkshop);
            firstProblemStatement.Workshop = firstWorkshop;
            firstWorkshop.Content.Add(firstProblemStatement);

            modelBuilder.Entity<Workshop>().HasData(firstWorkshop);
            modelBuilder.Entity<PropblemStatement>().HasData(firstProblemStatement);
            modelBuilder.Entity<Tag>().HasData(firstTag);

            #endregion

            base.OnModelCreating(modelBuilder);

        }
    }
}
