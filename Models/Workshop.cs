using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;

namespace estuary_backend.Models
{
    public class Workshop
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id {  get; set; }

        public DateTime Date { get; set; }

        public string Teaser { get; set; }

        public string LocationName { get; set; }

        public string LocationMap {  get; set; }

        public ICollection<Tag> Tags { get; set; }

        public List<Content> Content { get; set; }

        public List<Author> Authors {  get; set;}

    }
}
