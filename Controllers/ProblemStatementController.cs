using estuary_backend.Models;
using estuary_backend.Services;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Collections.Generic;
using System.Linq;

namespace estuary_backend.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class ProblemStatementController : ControllerBase
    {
        private readonly ILogger<ProblemStatementController> logger;
        private readonly IWorkshopService workshopService;

        public ProblemStatementController(ILogger<ProblemStatementController> logger, IWorkshopService workshopService)
        {
            this.logger = logger;
            this.workshopService = workshopService;
        }
        // GET: api/<ProblemStatementController>
        [HttpPost]
        public IEnumerable<PropblemStatement> Post([FromBody] int workshopId)
        {
            using var ctx = new EstuaryDbContext();
            var propblemStatements = ctx.Contents
                .Where(content => content.Workshop.Id == workshopId)
                .Where(content => content is PropblemStatement)
                .Select(content => content as PropblemStatement)
                .ToList();

            this.logger.LogInformation($"Get all Problemstatements ({propblemStatements.Count}) for workshop with id '{workshopId}'");

            return propblemStatements;
        }

        [HttpPost("{workshopId}")]
        public IActionResult AddProblemStatement(int workshopId, [FromBody] PropblemStatement propblemStatement)
        {
            if (!workshopService.WorkshopExists(workshopId))
            {
                this.logger.LogInformation($"Unable to add Problem Statement to Workshop. Workshop with id '{workshopId}' not found.");
                return NotFound();
            }

            using var ctx = new EstuaryDbContext();
            var workshop = ctx.Workshops.Where(workshop => workshop.Id == workshopId).FirstOrDefault();

            propblemStatement.Id = 0;
            ctx.Add(propblemStatement);
            if (workshop.Content == null)
            {
                workshop.Content = new List<Content>();
            }
            workshop.Content.Add(propblemStatement);

            ctx.SaveChanges();

            this.logger.LogInformation($"Added new Problem Statement '{propblemStatement.Id}' to Workshop '{workshopId}'");
            return Ok();
        }

    }
}
