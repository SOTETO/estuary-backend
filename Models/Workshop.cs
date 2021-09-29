using System;
using System.Collections.Generic;

namespace estuary_backend.Models
{
    public class Workshop
    {
        public Guid Id {  get; set; }

        public DateTime Date { get; set; }

        public string Teaser { get; set; }

        public Location Location { get; set; }

        public List<Tag> Tags { get; set; }

        public List<Content> Content { get; set; } = new List<Content>();

        public List<Author> Authors {  get; set;} = new List<Author>();

    }
}
