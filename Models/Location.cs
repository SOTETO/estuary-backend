using System;
using System.ComponentModel.DataAnnotations.Schema;

namespace estuary_backend.Models
{
    public class Location
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id { get; set; }
        public string Name { get; set; }

        public string Map {  get; set; }    
    }
}
