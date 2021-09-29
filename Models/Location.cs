using System;

namespace estuary_backend.Models
{
    public class Location
    {
        public Guid Id { get; set; }
        public string Name { get; set; }

        public string Map {  get; set; }    
    }
}
