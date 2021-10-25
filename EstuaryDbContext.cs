using estuary_backend.Models;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Diagnostics;
using System;
using System.Collections.Generic;

namespace estuary_backend
{
    public class EstuaryDbContext : DbContext
    {
        protected override void OnConfiguring(DbContextOptionsBuilder options) 
            => options
                .UseSqlite("Data Source=estuary_backend.db")
                .LogTo(Console.WriteLine, new[] { RelationalEventId.CommandExecuted })
                .EnableSensitiveDataLogging();

        public DbSet<Workshop> Workshops { get; set; }

        public DbSet<Content> Contents { get; set; }

        public DbSet<Tag> Tags { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<PropblemStatement>();
        }
/*  
        protected override async void OnModelCreating(ModelBuilder modelBuilder)
        {
          
            #region seed test data

            /*

            var firstProblemStatement = new PropblemStatement
            {
                //WorkshopId = firstWorkshop.Id,
                Title = "ProblemStatement 1",
                Iam = "Lorem ipsum",
                Iwant = "dolor sit amet, consetetur sadipscing elitr",
                But = "sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat",
                Because = "sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.",
                Feel = "Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.Lorem ipsum dolor sit amet, consetetur sadipscing."
            };

            modelBuilder.Entity<PropblemStatement>().HasData(firstProblemStatement);


            // var workshopTag = new WorkshopTag { Id = Guid.NewGuid(), TagId = firstTag.Id, WorkshopId = firstWorkshop.Id };

            modelBuilder.Entity<WorkshopTag>().HasData(workshopTag);
            

            #endregion
         
               
            base.OnModelCreating(modelBuilder);

        }*/
    }
}
