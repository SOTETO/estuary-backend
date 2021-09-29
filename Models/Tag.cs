using System;
using System.Collections.Generic;

namespace estuary_backend.Models
{
    public class Tag
    {
        public Guid Id { get; set; }

        public string Name { get; set; }

        public List<WorkshopTag> WorkshopTags { get; set; }
    }
}
