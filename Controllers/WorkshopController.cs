using estuary_backend.Models;
using Microsoft.AspNetCore.Mvc;
using System.Collections.Generic;

// For more information on enabling Web API for empty projects, visit https://go.microsoft.com/fwlink/?LinkID=397860

namespace estuary_backend.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class WorkshopController : ControllerBase
    {
        // GET: api/<WorkshopController>
        [HttpGet]
        public IEnumerable<Workshop> Get()
        {
            return new Workshop[] { new Workshop(), new Workshop() };
        }

        // GET api/<WorkshopController>/5
        [HttpGet("{id}")]
        public string Get(int id)
        {
            return $"value: {id}";
        }

        // POST api/<WorkshopController>
        [HttpPost]
        public string Post([FromBody] string value)
        {
            return $"Post {value}";
        }

        // PUT api/<WorkshopController>/5
        [HttpPut("{id}")]
        public string Put(int id, [FromBody] string value)
        {
            return $"Put {value} to {id}";
        }

        // DELETE api/<WorkshopController>/5
        [HttpDelete("{id}")]
        public string Delete(int id)
        {
            return $"Delete: {id}";
        }
    }
}
