using System;
using System.Collections.Generic;

namespace estuary_backend.Models
{
    public class Tag
    {
        public string Name { get; set; }

        public List<Workshop> Workshops { get; set; } = new List<Workshop>();
    }
}
