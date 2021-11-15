using estuary_backend.DTOs;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;

namespace estuary_backend.Models
{
    public class Workshop
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id { get; set; }

        public DateTime Date { get; set; }

        public string Teaser { get; set; }

        public string LocationName { get; set; }

        public string LocationMap { get; set; }

        public ICollection<Tag> Tags { get; set; }

        public List<Content> Content { get; set; }

        public List<Author> Authors { get; set; }

        public void Update(Workshop source)
        {
            this.Date = source.Date;
            this.Teaser = source.Teaser;
            this.LocationName = source.LocationName;
            this.LocationMap = source.LocationMap;
            this.Tags = source.Tags;
            this.Content = source.Content;
            this.Authors = source.Authors;
        }

        public WorkshopDTO toDTO()
        {
            return new WorkshopDTO
            {
                id = Id,
                type = 0, // TODO
                locationName = LocationName,
                locationMap = LocationMap,
                date = Date.Date,
                tags = Tags.Select(tag => tag.Name).ToArray(),
                likes = Content.Select(content => content.Likes.Count).Sum(),
                teaser = Teaser,
                authors = Authors.Where(author => author.Visible).Select(author => author.DropsID).ToArray(),
                status = 0,
                content = Content
                    .Where(content => content is ProblemStatement)
                    .Select(content => ((ProblemStatement)content).toDTO()).ToArray(),
            };
        }
    }
}
