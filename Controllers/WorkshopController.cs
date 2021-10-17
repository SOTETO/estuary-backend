using estuary_backend.Models;
using estuary_backend.Services;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;


namespace estuary_backend.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class WorkshopController : ControllerBase
    {
        private readonly IWorkshopService workshopService;

        public WorkshopController(IWorkshopService workshopService)
        {
            this.workshopService = workshopService;
        }

        // GET: api/<WorkshopController>
        [HttpGet]
        public IEnumerable<Workshop> Get()
        {
            using var ctx = new EstuaryDbContext();
            var workshops = ctx.Workshops
                .Include(ws => ws.Tags)
                .Include(ws => ws.Content)
                .Include(ws => ws.Authors)
                .ToList();

            return workshops;
        }

        // GET api/<WorkshopController>/5
        [HttpGet("{id}")]
        public Workshop Get(int id)
        {
            return workshopService.GetWorkshopById(id);
        }

        // GET api/<WorkshopController>/Tag:Improvements
        [HttpGet("Tag:{tag}")]
        public IEnumerable<Workshop> Get(string tag)
        {
            using var ctx = new EstuaryDbContext();
            var workshops = ctx.Workshops
                .Where(ws => ws.Tags.Any(t => string.Equals(t.Name, tag)))
                .Include(ws => ws.Tags)
                .Include(ws => ws.Content)
                .Include(ws => ws.Authors)
                .ToList();

            return workshops;
        }

        // POST api/<WorkshopController>
        [HttpPost]
        public int Post([FromBody] Workshop value)
        {
            return workshopService.CreateWorkshop(value);
        }

        // PUT api/<WorkshopController>/5
        [HttpPut("{id}")]
        public string Put(int id, [FromBody] string value)
        {
            return $"Put {value} to {id}";
        }

        // DELETE api/<WorkshopController>/5
        [HttpDelete("{id}")]
        public void Delete(int id)
        {
            workshopService.DeleteWorkshop(id);
        }

        #region for testing only

        // GET api/<WorkshopController>/Seed
        [HttpGet("Seed")]
        public string Seed()
        {
            using var ctx = new EstuaryDbContext();

            var tag1 = new Tag { Name = "Improvements" };
            var tag2 = new Tag { Name = "Tag2" };

            var workshop1 = new Workshop
            {
                Date = DateTime.Now,
                Authors = new List<Author> { },
                Teaser = "Lorem ipsum dolor sit amet.",
                LocationName = "Hamburg",
                LocationMap = "https://goo.gl/maps/mbnen1jr8C81J6vU9",
                Tags = new List<Tag> { tag1, tag2 }
            };

            var workshop2 = new Workshop
            {
                Date = DateTime.Now,
                Authors = new List<Author> { },
                Teaser = "Lorem ipsum dolor sit amet.",
                LocationName = "Berlin",
                LocationMap = "https://goo.gl/maps/mbnen1jr6C21J6vU9",
                Tags = new List<Tag> { tag1 }
            };

            ctx.AddRange(tag1, tag2, workshop1, workshop2);

            ctx.SaveChanges();

            return "OK";
        }

        #endregion
    }
}
