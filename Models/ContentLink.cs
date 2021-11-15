using estuary_backend.DTOs;
using System;
using System.ComponentModel.DataAnnotations.Schema;

namespace estuary_backend.Models
{
    public class ContentLink
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id { get; set; }

        public Content Content { get; set; }

        public string LinkTag { get; set; }

        public ContentLinkDTO toDTO()
        {
            return new ContentLinkDTO
            {
                id = Content.Id,
                tag = LinkTag,
            };
        }
    }
}
